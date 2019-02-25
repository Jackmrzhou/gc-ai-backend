#define LINUX
#include<fstream>
#include<iostream>
#include<utility>
#include<vector>
#include<string>
#include<stdlib.h>
#include<cstdio>
#include<ctime>
#ifdef LINUX
#include<unistd.h>
#include<sys/wait.h>
#else
#include<windows.h>
#endif
using namespace std;
//const string textA = string{ "playerA.cpp" };
//const string programA = string{ "playerA" };
//const string textB = string{ "playerB.cpp" };
//const string programB = string{ "playerB" };
//const string judgment = string{ "result.txt" };
const int waitTime = 1000;//the wait time of each turn
void updateInput(ofstream&out, int &rounds, vector<pair<int, int> >& steps) {
	out.open("input.txt");
	out << rounds << endl;
	for (int i = 0; i < rounds; ++i) {
		out << steps[i].first << " " << steps[i].second << endl;
	}
	out.close();
}
#ifdef LINUX
int executePackage(const string&processName){
	//0  -->normal
	//-1 -->time out
	//-2 -->runing error
	pid_t fpid = fork();
	if(fpid == 0){//son process
		if(execl(processName.c_str(),"", NULL)<0)
			exit(-1);
		exit(0);
	}
	else{
		clock_t end = clock() + waitTime,now;
		int status,ret;
		do {
			ret = waitpid(fpid, &status, WNOHANG);
		} while (ret == 0 && end >= (now=clock()));
		if(WIFEXITED(status))
			if(WEXITSTATUS(status)	== -1)
				return -2;
			else
				return 0;
		if(end < now)
			return -1;
		return -2;
	}
	
}
#else
LPCWSTR stringToLPCWSTR(const string&orig) {
	wchar_t *wcstring = 0;
	try {
		size_t origsize = orig.length() + 1;
		const size_t newsize = 100;
		size_t convertedChars = 0;
		if (orig == "") {
			wcstring = (wchar_t *)malloc(0);
			mbstowcs_s(&convertedChars, wcstring, origsize, orig.c_str(), _TRUNCATE);
		}
		else {
			wcstring = (wchar_t *)malloc(sizeof(wchar_t)*(orig.length() - 1));
			mbstowcs_s(&convertedChars, wcstring, origsize, orig.c_str(), _TRUNCATE);
		}
	}
	catch (std::exception e) {}
	return wcstring;
}
int executePackage(LPCWSTR fileName, LPCWSTR args, LPCWSTR dir)
{
	SHELLEXECUTEINFOW sei = { sizeof(SHELLEXECUTEINFOW) };

	sei.fMask = SEE_MASK_NOCLOSEPROCESS | SEE_MASK_FLAG_NO_UI;

	sei.lpFile = fileName;
	sei.lpParameters = args;
	sei.lpDirectory = dir;

	if (!ShellExecuteExW(&sei)) {
		return false;
	}

	HANDLE hProcess = sei.hProcess;
	int res = -1;
	if (hProcess != 0) {
		res = WaitForSingleObject(hProcess, waitTime);
		CloseHandle(hProcess);
	}

	if (res == WAIT_TIMEOUT)
		return -1;
	if (res == WAIT_OBJECT_0)
		return 0;
	return -2;
}
#endif
inline int myAbs(const int&t) {
	return t < 0 ? -t: t;
}

/**
 *tic-tac-toe judger
 *suppose the x go first
 *
 *argv={
 *	program name,
 *	result file position,("result.txt")
 *	playerA file name,("playerA.cpp")
 *	playerA exe name,("playerA")
 *	playerB file name,("playerB.cpp")
 *	playerB exe name("playerB")
 *}
 *
 */
int main(int argc, char *argv[]) {
	const string judgement{ argv[1] };
	const string textA{ argv[2] };
	const string programA{ argv[3] };
	const string textB{ argv[4] };
	const string programB{ argv[5] };

	ofstream result{ judgement };

#ifdef LINUX
	system(("rm " + programA).c_str());
	system(("rm " + programB).c_str());
	system(("g++ " + textA + " -o " + programA).c_str());
	system(("g++ " + textB + " -o " + programB).c_str());
	int resA = access((programA).c_str(), F_OK);
	int resB = access((programB).c_str(), F_OK);
	
	bool flag = false;
	if (resA != 0) {
		result << "CompileErrorA" << " ";
		flag = true;
	}
	if (resB != 0) {
		result << "CompileErrorB";
		flag = true;
	}
	if (flag)
		return 0;
#else
	system(("erase " + programA + ".exe").c_str());
	system(("erase " + programB + ".exe").c_str());
	system(("g++ " + textA + " -o " + programA).c_str());
	system(("g++ " + textB + " -o " + programB).c_str());
	int resA = WinExec((programA + ".exe").c_str(), SW_SHOWNORMAL);
	int resB = WinExec((programB + ".exe").c_str(), SW_SHOWNORMAL);

	bool flag = false;
	//printf("%d %d \n",resA,resB);//********************************
	if (resA == ERROR_FILE_NOT_FOUND) {
		result << "CompileErrorA" << " ";
		flag = true;
	}
	if (resB == ERROR_FILE_NOT_FOUND) {
		result << "CompileErrorB";
		flag = true;
	}
	if (flag)
		return 0;
#endif

	ofstream out;
	ifstream in;
	int rounds = 0;
	vector<pair<int, int> >steps;
#ifdef LINUX
	const string programs[2] = {programA,programB};
#else
	const LPCWSTR programs[2] = { stringToLPCWSTR(programA),stringToLPCWSTR(programB) };
	const LPCWSTR arg = stringToLPCWSTR("");
	const LPCWSTR dir = stringToLPCWSTR(".\\");
#endif
	int winner = -1;//1->A 0->B -1->tie
	int row[3] = {}, col[3] = {}, cross[2] = {};
	while (rounds < 9) {
		//printf("%d\n", rounds);//****************************
		updateInput(out, rounds, steps);
		++rounds;

		int res =//run the program
#ifdef LINUX
			executePackage(programs[rounds & 1 ^ 1]);
#else
			executePackage(programs[rounds & 1 ^ 1], arg, dir);
#endif

		int x = -1, y = -1, cache;
		if (res == -1)
			y = -3;//-1 -3-->out of time
		else if (res == -2)
			y = -2;//-1 -2-->run time error
		else {
			in.open("output.txt");
			while (!in.eof()) {
				in >> x >> y >> cache;
			}
			in.close();

			if (x == -1 || y == -1) {//-1 -1-->wrong format
				x = -1;
				y = -1;
			}
		}

		steps.push_back(pair<int, int>{x, y});
		if (x == -1) {
			winner = 1 ^ (rounds & 1);//the player fails
			break;
		}

		int add = (rounds & 1) == 0 ? 1 : -1;
		row[x] += add;
		col[y] += add;
		if (x == y)
			cross[0] += add;
		if (x + y == 2)
			cross[1] += add;
		if (myAbs(row[x]) == 3 || myAbs(col[y]) == 3 || myAbs(cross[0]) == 3 || myAbs(cross[1]) == 3) {
			winner = (rounds & 1);//the player wins
			break;
		}
	}

	result << rounds << " ";
	for (int i = 0; i < rounds; ++i)
		result << steps[i].first << " " << steps[i].second << " ";
	char charWinner;
	if (winner == -1)
		charWinner = 'T';
	else if (winner == 1)
		charWinner = 'A';
	else
		charWinner = 'B';
	result << charWinner;

	return 0;

}
