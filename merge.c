# include <stdio.h>

int main( int argc, char *argv[]){
    
    if( argc == 2){
        printf("The argument supplied is %s\n", argv[1]);
    }

    else if(argc > 2){
        printf("To many arguments supplied.\n");
    }

    else {
        printf("One argument is expected.\n");
    }


}
