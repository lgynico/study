#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

void* threadA_func(void *);
void* threadB_func(void *);

pthread_mutex_t mutexA = PTHREAD_MUTEX_INITIALIZER;
pthread_mutex_t mutexB = PTHREAD_MUTEX_INITIALIZER;

int main(void)
{

    pthread_t threadA;
    pthread_t threadB;

    pthread_create(&threadA, NULL, threadA_func, NULL);
    pthread_create(&threadB, NULL, threadB_func, NULL);

    pthread_join(threadA, NULL);
    pthread_join(threadB, NULL);

    printf("exit\n");

    return;
}

void* threadA_func(void *)
{
    printf("Thread A wait for Resource A...\n");
    pthread_mutex_lock(&mutexA);
    printf("Thread A got Resource A.\n");

    sleep(1);

    printf("Thread A wait for Resource B...\n");
    pthread_mutex_lock(&mutexB);
    printf("Thread A got Resource B.\n");


    pthread_mutex_unlock(&mutexB);
    pthread_mutex_unlock(&mutexA);

    printf("Thread A done!\n");

    return 0;
}

void* threadB_func(void *)
{
    printf("Thread B wait for Resource B...\n");
    pthread_mutex_lock(&mutexB);
    printf("Thread B got Resource B.\n");

    sleep(1);

    printf("Thread B wait for Resource A...\n");
    pthread_mutex_lock(&mutexA);
    printf("Thread B got Resource A.\n");


    pthread_mutex_unlock(&mutexA);
    pthread_mutex_unlock(&mutexB);

    printf("Thread B done!\n");

    return 0;
}