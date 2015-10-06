
#include "cuda_runtime.h"
#include "device_launch_parameters.h"
#include "C:/Users/Lin/Documents/visual studio 2012/Projects/common/book.h"


#include <stdio.h>

__global__ void add( int a, int b, int *c ) {
    *c = a + b;
	printf("%d", a);
}

int main( void ) {
    int c;
    int *dev_c;
    HANDLE_ERROR( cudaMalloc( (void**)&dev_c, sizeof(int) ) );

    add<<<1,1>>>( 2, 7, dev_c );

    HANDLE_ERROR( cudaMemcpy( &c, dev_c, sizeof(int),
                              cudaMemcpyDeviceToHost ) );
    printf( "2 + 7 = %d\n", c );
    HANDLE_ERROR( cudaFree( dev_c ) );

    return 0;
}
