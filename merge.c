# include <stdio.h>
# include <dirent.h>

int read(void){
    
    DIR *d;
    struct dirent *dir;
    d = opendir(".");
    if(d){
        while((dir = readdir(d)) != NULL){
            printf("%s\n", dir->d_name);
        }
        closedir(d);

    }
    return(0);
}


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

    read();


}
