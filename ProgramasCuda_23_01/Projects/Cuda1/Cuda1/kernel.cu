#include "cuda_runtime.h"
#include "device_launch_parameters.h"
#include <stdio.h>
#include <windows.h>

const int arraySize = 501;

void somaVetorCuda(const int *a, const int *b, int *c, unsigned int size);

void somaVetorLoop(int a[arraySize], int b[arraySize], int c[arraySize], int arraySize);

__global__ void somaKernel(int *c, const int *a, const int *b)
{
    int i = threadIdx.x;
    c[i] = a[i] + b[i];
}

int main()
{
	int inicio, final; 


	int a[arraySize];
	int b[arraySize];
	int c[arraySize];

	for (int i=0; i<arraySize;i++){
		a[i] = i;
		b[i] = i;
	}

    inicio = GetTickCount();
    // Add vectors in parallel.
    //somaVetorCuda(a, b, c, arraySize);
	somaVetorLoop(a, b, c, arraySize);

	final = GetTickCount();
    printf("\n\ntempo decorrido: %d\n", final-inicio);
	for (int i=0;i<arraySize;i++){
		printf("%d, ", c[i]);
	}

	

	

    // cudaDeviceReset must be called before exiting in order for profiling and
    // tracing tools such as Nsight and Visual Profiler to show complete traces.
    cudaDeviceReset();

    return 0;
}

void somaVetorLoop(int a[arraySize], int b[arraySize], int c[arraySize], int arraySize){
	for (int i=0; i<arraySize;i++){
		c[i] = a[i] + b[i];
	}
}

// Função que soma os vetores de modo paralelo
void somaVetorCuda(const int *a, const int *b, int *c, unsigned int size)
{
    int *dev_a = 0;
    int *dev_b = 0;
    int *dev_c = 0;
   
    // Allocate GPU buffers for three vectors (two input, one output)
    cudaMalloc((void**)&dev_c, size * sizeof(int));
    cudaMalloc((void**)&dev_a, size * sizeof(int));
    cudaMalloc((void**)&dev_b, size * sizeof(int));

    // Copy input vectors from host memory to GPU buffers.
    cudaMemcpy(dev_a, a, size * sizeof(int), cudaMemcpyHostToDevice);
    cudaMemcpy(dev_b, b, size * sizeof(int), cudaMemcpyHostToDevice);

    // Launch a kernel on the GPU with one thread for each element.
    somaKernel<<<1, size>>>(dev_c, dev_a, dev_b);
    
    // cudaDeviceSynchronize waits for the kernel to finish
    //cudaDeviceSynchronize();

    // Copy output vector from GPU buffer to host memory.
    cudaMemcpy(c, dev_c, size * sizeof(int), cudaMemcpyDeviceToHost);

}
