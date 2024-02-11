
function lengt(figura) {
    let max = 0
    figura = figura.toString().split("\n")
    for(let i = 0; i < figura.length;i++) {
        if (max < figura[i].length) {
            max = figura[i].length
        }
    }
    return max
}

function height (figura) {
    let max = 0
    figura = figura.toString().split("\n")
    max = figura.length
    return max
}

function argument1() {
    let num = 0

    num = process.argv[2]

    return num
}

function argument2() {
    file = process.argv[3]
    let fs = require("fs");
    let data = fs.readFileSync(file);

    let figura = data.toString()

    return figura
}

function createarray(y, x) {
    let array = []
    for (let i = 0; i < y; i++) {
        console.log()
        for(let j = 0; j < x; j++){
            array.push(0)
        }
    }
    return array
}

function readfile(file) {
    let x = lengt(file)
    let y = height(file)
    let array = createarray(y, x)
    array = add_value(array, file, x)

    return array
}

function add_value(array, file, x) {
    let figura = file.split("\n")
    for (let i = 0;i < figura.length; i++) {
        let len = figura[i]
        for (let j = 0; j < len.length; j++) {
            if (len[j] == "*") {
                array[i * x + j] = 1
            }
        }
    }
    return array
}

function field (array, y, x) {
    for (let i = 0; i < y; i++) {
        console.log()
        for (let j = 0; j < x; j++) {
            if (array[i*x+j] == 0) {
                process.stdout.write("-")
            } else {
                process.stdout.write("*")
            }
        }
    }
}

function livedead(array, osy, osx) {
	let new_array = clone(array)
	for (y = 0; y < osy; y++) {
		for (x = 0; x < osx; x++) {
			let counter = cell_counter(array, y, x, osx, osy)
			if ((counter > 3 || counter < 2) && array[y*osx+x] == 1) {
				new_array[y*osx+x] = 0
			}
			if (counter == 3 && array[y*osx+x] == 0) {
				new_array[y*osx+x] = 1
			}
		}
	}
	return new_array
}

function cell_counter(array , y, x, osx, osy) {
	let res = 0
	res += predel(array, y, x+1, osx, osy)
	res += predel(array, y, x-1, osx, osy)
	res += predel(array, y+1, x, osx, osy)
	res += predel(array, y-1, x, osx, osy)
	res += predel(array, y-1, x-1, osx, osy)
	res += predel(array, y-1, x+1, osx, osy)
	res += predel(array, y+1, x-1, osx, osy)
	res += predel(array, y+1, x+1, osx, osy)
	return res
}
function predel(array, y, x, osx, osy) {
	if ((y >= 0 && x >= 0) && (y < osy && x < osx)) {
		if (array[y*osx+x] == 1) {
			return 1
		}
	}
	return 0
}
function clone (src) {
    return JSON.parse(JSON.stringify(src));
}

function gLife(array , cycle_num, y, x ) {
	for (let i = 0; i < cycle_num; i++) {
		field(array, y, x)
		array = livedead(array, y, x)
	}

}
//
let str = argument1()
let figura = argument2()
//
let num = Number(str) //arg1
//
let array = readfile(figura)
let x = lengt(figura)
let y = height(figura)
gLife(array,num,y,x)
//

