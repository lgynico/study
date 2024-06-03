#include <stdio.h>
#include <stdlib.h>
#include "stackADT.h"

struct node
{
    Item data;
    struct node *next;
};

struct stack_type
{
    struct node *top;
};

static void ternimate(const char *message)
{
    printf("%s\n", message);
    exit(EXIT_FAILURE);
}

Stack create()
{
    Stack s = malloc(sizeof(struct stack_type));
    if (s == NULL)
    {
        ternimate("Error in create: stack could not be created.");
    }

    s->top = NULL;
    return s;
}

void destroy(Stack s)
{
    make_empty(s);
    free(s);
}

void make_empty(Stack s)
{
    while (!is_empty(s))
    {
        pop(s);
    }
}

bool is_empty(Stack s)
{
    return s->top == NULL;
}

bool is_full(Stack s)
{
    return false;
}

void push(Stack s, Item i)
{
    if (is_full(s))
    {
        ternimate("Error in push: stack is full.");
    }

    struct node *new_node = malloc(sizeof(struct node));
    if (new_node == NULL)
    {
        ternimate("Error in push: stack is full");
    }

    new_node->data = i;
    new_node->next = s->top;
    s->top = new_node;
}

Item pop(Stack s)
{
    if (is_empty(s))
    {
        ternimate("Error in pop: stack is empty.");
    }

    struct node *old_top = s->top;
    Item i = old_top->data;
    s->top = old_top->next;

    free(old_top);

    return i;
}