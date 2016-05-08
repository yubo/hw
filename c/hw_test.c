// Copyright 2016, yubo.
// All rights reserved.

#include "hw.h"

#define BSIZE 64

char buff[BSIZE+1];

static int testHw(void){
	const char *want = "hello,world";

	hw(buff, BSIZE);
	if (strncmp(want, buff, BSIZE)){
		printf("got \"%s\" want \"%s\"\n", buff, want);
		return -1;
	}
	return 0;
}

static int testHwi(void){
	size_t i;

	struct {
		int n;
		char want[BSIZE+1];
	} tests[] = {
		{5, "hello,world small"},
		{50, "hello,world big"},
	};

	for(i = 0; i < sizeof(tests)/sizeof(tests[0]); i++){
		hwi(tests[i].n, buff, BSIZE);
		if (strncmp(tests[i].want, buff, BSIZE)){
			printf("got \"%s\" want \"%s\"\n", buff, tests[i].want);
			return -1;
		}
	}
	return 0;
}

int main(){

	if(testHw())
		return -1;

	if(testHwi())
		return -1;

	printf("PASS\n");
	return 0;
}
