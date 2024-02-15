import ctypes
import sys
from ctypes import *

count = CDLL('./count.so')
 
def clone(array) : 
    new_array = array.copy()
    return new_array

def argument1() :
    number = int(sys.argv[1])

    return number

def argument2() :
    filename = sys.argv[2]
    with open(filename,'r') as file:
        lines = file.readlines()

    return lines

def lenght(file) :
    max = 0

    for i in range(len(file)):
        if max < len(file[i]) :
            max = len(file[i])

    return max
        
def height(file) :
    max = len(file)
    return max

def createarr(y, x) :
    arr = []
    for i in range(y):
        for j in range(x):
            arr.append(0)
    
    return arr

def addvalue(array, x, file) :
    for i in range(len(file)) :
        line = file[i]
        for j in range (len(line)) :
            if line[j] == "*" :
                array [i * x + j] = 1
    
    return array

def livedead(array, osy, osx) :
    new_array = clone(array)
    
    
    arrayc = (ctypes.c_byte * len(array))(*array)
    
    
    for y in range(osy) :
        for x in range (osx) :
            counter =  count.cell_counter(arrayc, y, x, osx, osy)
            if ((counter > 3 or counter < 2) and array[y*osx+x] == 1) :
                new_array[y*osx+x] = 0
			
            if (counter == 3 and array[y*osx+x] == 0) :
                 new_array[y*osx+x] = 1
			
		
	
    return new_array


def readfile(file) :
    x = lenght(file)
    y = height(file)

    array = createarr(y, x)
    array = addvalue(array, x, file)
    return array

def field(array, y, x) :
    for i in range(y):
        print("\n")
        for j in range(x):
            if array[i * x + j] == 0 :
                print("-",end=' ')
            else :
                print("*",end=' ')

def gLife(array , cycle_num, y, x ) :
	for i in range(cycle_num) :
		field(array, y, x)
		array = livedead(array, y, x)
	



num = argument1()
file = argument2()
array = readfile(file)
x = lenght(file)
y = height(file)
gLife(array,num,y,x)