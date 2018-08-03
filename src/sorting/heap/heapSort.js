'use strict';

var heapSort = (function () {

  function swap (array, i, j) {
    var temp = array[i];
    array[i] = array[j];
    array[j] = temp;
    return array;
  }

  function heapify(array, index, heapSize) {
    var left = index * 2 + 1;
    var right = index * 2 + 2;
    var largest = index;

    if (left < heapSize && array[left] > array[index]) {
      largest = left;
    }

    if (right < heapSize && array[right] > array[largest]) {
      largest = right;
    }

    if (largest !== index) {
      swap(array, index, largest);
      heapify(array, largest, heapSize);
    }
  }

  function buildMaxHeap(array) {
    for (var i = Math.floor(array.length / 2) - 1; i >= 0; i--) {
      heapify(array, i, array.length);
    }

    return array;
  }

  return function (array) {
    var size = array.length;

    buildMaxHeap(array);

    for (var i = array.length - 1; i > 0; i--) {
      swap(array, 0, i);
      size -= 1;
      heapify(array, 0, size);
    }

    return array;
  }
}());

module.exports = heapSort;