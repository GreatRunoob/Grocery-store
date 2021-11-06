#include<stdlib.h>
#include"stack.h"

_Bool stack_is_invalid(stack_ptr stack)
{
	return !stack_is_valid(stack);
};

snode_ptr init_snode(stck_elem_t elem)
{
	snode_ptr snode=(snode_ptr)malloc(sizeof(struct snode));
	if(snode)
	{
		snode->elem=elem;
		snode->next=NULL;
	}
	return snode;
}

stack_ptr init_stack()
{
	stack_ptr stack=(stack_ptr)malloc(sizeof(struct stack));
	if(stack)
	{
		stack->top=NULL;
		stack->size=STACK_EMPTY;
	}
	return stack;
}

stck_size_t get_stck_size(stack_ptr stack)
{
	return stack->size;
}

_Bool stack_is_empty(stack_ptr stack)
{
	return get_stck_size(stack)==STACK_EMPTY;
}

_Bool stack_is_valid(stack_ptr stack)
{
	return stack!=NULL;
}

_Bool push(stack_ptr stack, stck_elem_t elem)
{
	if(stack_is_invalid(stack))
		return FALSE;

	snode_ptr snode=init_snode(elem);
	if(!snode)
		return FALSE;

	snode->next=stack->top;
	stack->top=snode;
	stack->size++;
	return TRUE;
}

_Bool pop(stack_ptr stack, stck_elem_t* elem)
{
	if(stack_is_invalid(stack))
		return FALSE;

	if(stack_is_empty(stack))
		return FALSE;

	snode_ptr pop_node=stack->top;
	stack->top=pop_node->next;
	*elem=pop_node->elem;

	free(pop_node);
	stack->size--;
	return TRUE;
}

_Bool empty_stack(stack_ptr stack)
{
	if(stack_is_invalid(stack))
		return FALSE;

	stck_elem_t tmp;
	while(pop(stack,&tmp));
	return TRUE;
}

stack_ptr del_stack(stack_ptr stack)
{
	if(empty_stack(stack))
	{
		free(stack);
		stack=NULL;
	}
	return stack;
}

