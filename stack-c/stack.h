#ifndef _CSTACK_H_
#define _CSTACK_H_

typedef unsigned int stck_elem_t;
typedef unsigned int stck_size_t;

typedef struct snode
{
	stck_elem_t elem;
	struct snode* next;
} *snode_ptr;

typedef struct stack
{
	snode_ptr top;
	stck_size_t size;
} *stack_ptr;

snode_ptr init_snode(stck_elem_t elem);

stack_ptr init_stack();

stck_size_t get_stck_size(stack_ptr stack);

_Bool stack_is_empty(stack_ptr stack);

_Bool stack_is_valid(stack_ptr stack);

_Bool push(stack_ptr stack, stck_elem_t elem);

_Bool pop(stack_ptr stack, stck_elem_t* elem);

_Bool empty_stack(stack_ptr stack);

stack_ptr del_stack(stack_ptr stack);

#define STACK_EMPTY (stck_size_t) 0
#define FALSE (_Bool) 0
#define TRUE (_Bool) 1

#endif
