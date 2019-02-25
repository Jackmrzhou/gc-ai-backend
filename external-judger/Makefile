cc = g++
CXXFLAGS = -std=c++11
target = judger
source = judger2.1.cpp
obj = judger2.1.o

$(target): $(obj)
	$(cc) $(CXXFLAGS) -o $(target) $(obj)

%.o: %.cpp
	$(cc) $(CXXFLAGS) -c $< -o $@

clean:
	rm -rf $(obj) $(target)