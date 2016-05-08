// Copyright 2016, yubo.
// All rights reserved.

#include <stdio.h>

int hw(char *buff, size_t size)
{
	return snprintf(buff, size, "hello,world");
}

int hwi(int a, char *buff, size_t size)
{
	if (a < 0){
		return snprintf(buff, size, "hello,world negative");
	}else if (a == 0 ){
		return snprintf(buff, size, "hello,world zero");
	}else if (a < 10 ){
		return snprintf(buff, size, "hello,world small");
	}else if (a < 100 ){
		return snprintf(buff, size, "hello,world big");
	}else if (a < 1000 ){
		return snprintf(buff, size, "hello,world huge");
	}
	return snprintf(buff, size, "hello,world enormous");
}
