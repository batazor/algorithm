//
// Created by batazor on 21.06.17.
//

#include <iostream>

int Factorial(int x) {
  if (x <= 1) {
    return 1;
  } else {
    return x * Factorial(x - 1);  // вычисляем факториал от x-1 и умножаем на x
  }
}

int main() {
    cout << Factorial(1) << endl;
    cout << Factorial(-2) << endl;
    cout << Factorial(4) << endl;
    return 0;
}
