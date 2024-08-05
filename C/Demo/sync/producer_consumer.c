#include <stdio.h>
#include <semaphore.h>

#define N 100

sem_t mutex;
sem_t fullBuffers;
sem_t emptyBuffers;


void producer()
{
    while (1)
    {

    }
}


int main(void)
{
    sem_init(&mutex, 0, 1);
    sem_init(&fullBuffers, 1, 0);
    sem_init(&emptyBuffers, 1, N);
}