CXX_FLAGS := -Wall -Wextra -std=c++17 -O3 -fPIC

SRC		:= .
INCLUDE	:= ctbignum 
LIBS 	:= -Bstatic -lstdc++

all: library

library: $(SRC)/*.cpp
	$(CXX) $(CXX_FLAGS) -isystem$(INCLUDE) $^ $(LIBS) -shared -o libeip1962.so

