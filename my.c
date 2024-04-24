#include <stdio.h>


int arraySize = 10;

int main()
{
    struct MyStruct *array = myFunction();
    
    struct MyStruct **array = myFunction();

    

    return 0;
}

struct MyStruct
{
    int roll;
    char name[20];
    float marks;
};

struct MyStruct *myFunction()
{
    struct MyStruct *array = malloc(sizeof(struct MyStruct) * arraySize);
    // Initialize array elements

    return array;
}

struct MyStruct **myFunction()
{
    struct MyStruct **array = malloc(sizeof(struct MyStruct *) * arraySize);
    for (int i = 0; i < arraySize; i++)
    {
        array[i] = malloc(sizeof(struct MyStruct));
        // Initialize individual structs
    }
    return array;
}