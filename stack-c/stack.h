#ifndef _CSTACK_H_
#define _CSTACK_H_

#include <stdbool.h>
#include <stddef.h>

#define STACK_EMPTY (stack_size_t) 0

typedef unsigned int stack_elem_t;
typedef unsigned int stack_size_t;

typedef struct snode
{
	stack_elem_t elem;
	struct snode *next;
} *snode_ptr;

typedef struct stack
{
	snode_ptr top;
	stack_size_t size;
} *stack_ptr;

static inline stack_size_t get_stack_size(stack_ptr stack)
{
	return stack->size;
}

static inline bool stack_is_valid(stack_ptr stack)
{
	return stack != NULL;
}

static bool stack_is_empty(stack_ptr stack);

static snode_ptr init_snode(stack_elem_t elem);

stack_ptr init_stack();

bool push(stack_ptr stack, stack_elem_t elem);

bool pop(stack_ptr stack, stack_elem_t *elem);

bool empty_stack(stack_ptr stack);

stack_ptr del_stack(stack_ptr stack);

#endif
