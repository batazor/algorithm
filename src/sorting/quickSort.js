'use strict';

var quickSort = (function () {

  function swap (array, i, j) {
    var temp = array[i];
    array[i] = array[j];
    array[j] = temp;
    return array;
  }

  function partition (array, p, r) {
    var x = array[r - 1];
    var i = p;
    for (var j = p; j < r - 1; j++) {
      if (array[j] <= x) {
        swap(array, j, i);
        i++;
      }
    }
    swap(array, i, r - 1);
    return i;
  }

  function quickSort (array, p, r) {
    if (p < r) {
      var q = partition(array, p, r);
      quickSort(array, p, q);
      quickSort(array, q + 1, r);
    }

    return array;
  }

  return function (array) {
    return quickSort(array, 0, array.length);
  }
}());

module.exports = quickSort;