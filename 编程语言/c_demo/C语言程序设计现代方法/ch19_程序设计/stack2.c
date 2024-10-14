#include <stdio.h>
#include <stdlib.h>
#include "stack.h"

struct node
{
    int data;
    struct node *next;
};

static struct node *top = NULL;

static void terminate(const char *message)
{
    printf("%s\n", message);
    exit(EXIT_FAILURE);
}

void make_empty(void)
{
    top = NULL;
}

bool is_empty(void)
{
    return top == NULL;
}

bool is_full(void)
{
    return false;
}

void push(int i)
{
    struct node *new_node = malloc(sizeof(struct node));
    if (new_node == NULL)
    {
        terminate("Error in push: stack is full.");
    }

    new_node->data = i;
    new_node->next = top;
    top = new_node;
}

int pop(void)
{
    struct node *old_top;
    int i;

    if (is_empty())
    {
        terminate("Error in pop: stack is empty.");
    }

    old_top = top;
    top = old_top->next;
    i = old_top->data;

    free(old_top);

    return i;
}

int main(void)
{
    push(1);
    push(2);
    push(3);
    push(4);

    while (!is_empty())
    {
        printf("%d\n", pop());
    }

    return 0;
}