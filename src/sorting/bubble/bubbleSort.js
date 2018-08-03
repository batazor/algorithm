'use strict';

var bubbleSort = function (array) {

  function swap (array, i, j) {
    var temp = array[i];
    array[i] = array[j];
    array[j] = temp;
    return array;
  }

  for (var i = 0; i < array.length; i += 1) {
    for (var j = i; j > 0; j--) {
      if (array[j] < array[j - 1]) {
        swap(array, j, j - 1);
      }
    }
  }

  return array;
};

module.exports = bubbleSort;