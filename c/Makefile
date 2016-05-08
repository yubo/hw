.PHONY: clean test

all: libhw.so

#profile-arcs:  .gcda
#test-coverage: .gcno
CFLAGS := -fPIC -fprofile-arcs -ftest-coverage

.c.o:
	gcc -c $(CFLAGS) $< -o $@

libhw.so: hw.o
	gcc --share -o libhw.so hw.o -lgcov

test: libhw.so hw_test.o hw.h
	gcc -o hw_test hw_test.o -L. -lhw -lgcov
	LD_LIBRARY_PATH=. ./hw_test
	# text report
	gcov hw.c  #create hw.gcov from hw.gcno hw.gcda
	# html report
	lcov -d . -t hw_test -o hw_test.info -b . -c #create hw_test.info from *.gcda
	genhtml -o result hw_test.info

clean:
	rm -rf *.o hw_test *.c.gcov *.gcno *.gcda *.so *.info result
