#ifndef _CSTACK_H_
#define _CSTACK_H_

typedef unsigned int stck_elem_t;
typedef unsigned int stck_size_t;

typedef struct snode
{
	stck_elem_t elem;
	struct snode* next;
} *p_snode;

typedef struct stack
{
	p_snode top;
	stck_size_t size;
} *p_stack;

p_snode init_snode(stck_elem_t elem);

p_stack init_stack();

stck_size_t get_stck_size(p_stack stack);

_Bool stack_is_empty(p_stack stack);

_Bool stack_is_valid(p_stack stack);

_Bool push(p_stack stack, stck_elem_t elem);

_Bool pop(p_stack stack, stck_elem_t* elem);

p_stack del_stack(p_stack stack);

#define STACK_EMPTY (stck_size_t) 0
#define FALSE (_Bool) 0
#define TRUE (_Bool) 1

#endif
