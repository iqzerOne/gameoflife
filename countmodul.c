#include <stdio.h>
int predel(char* array,int y,int x,int osx,int osy);
int cell_counter(char* array,int y,int x,int osx,int osy);

int predel(char* array,int y,int x,int osx,int osy) {
	if ((y >= 0 && x >= 0) && (y < osy && x < osx)) {
		if (array[y*osx+x] == 1) {
			return 1;
		}
	}
	return 0;
}

int cell_counter(char* array,int y,int x,int osx,int osy) {
	int res = 0;
	res += predel(array, y, x+1, osx, osy);
	res += predel(array, y, x-1, osx, osy);
	res += predel(array, y+1, x, osx, osy);
	res += predel(array, y-1, x, osx, osy);
	res += predel(array, y-1, x-1, osx, osy);
	res += predel(array, y-1, x+1, osx, osy);
	res += predel(array, y+1, x-1, osx, osy);
	res += predel(array, y+1, x+1, osx, osy);
	return res;
}
