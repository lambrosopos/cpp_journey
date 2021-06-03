#include <iostream>
using namespace std;

/* 'this' pointer in c++
 *
 * 1. Each object gets its own copy of the data member.
 * 2. All-access the same function definition as present in the code segment
 *
 * The 'this' pointer is passed as a hidden argument to all nonstatic member function calls
 * and is available as local variable within the body of all nonstatic functions.
 * 
 * 'this' pointer is not available in static member functions as static member functions
 * can be called without any object (with class name)
 *
 * ex.) For a class X, the type of this pointer is 'X*'
 *
 * Also if a member function of X is declared as const, then the type of this pointer is
 * 'const X*'
 *
 * You can destroy an C++ object by calling 'delete this;'
 */

// Use case 1: local variable's name same as member's name

class Case1
{
    private:
        int x = 30;
    public:
        void setX (int x)
        {
            //this->x = x;
        }
        void print() { cout << "x = " << x << endl; }
};

int main()
{
    Case1 obj;
    int x = 20;
    obj.setX(x);
    obj.print();
    return 0;
}
