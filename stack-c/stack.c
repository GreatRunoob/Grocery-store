#include<stdlib.h>
#include"stack.h"

_Bool stack_is_invalid(p_stack stack)
{
	return !stack_is_valid(stack);
};

p_snode init_snode(stck_elem_t elem)
{
	p_snode snode=(p_snode)malloc(sizeof(struct snode));
	if(snode)
	{
		snode->elem=elem;
		snode->next=NULL;
	}
	return snode;
}

p_stack init_stack()
{
	p_stack stack=(p_stack)malloc(sizeof(struct stack));
	if(stack)
	{
		stack->top=NULL;
		stack->size=STACK_EMPTY;
	}
	return stack;
}

stck_size_t get_stck_size(p_stack stack)
{
	return stack->size;
}

_Bool stack_is_empty(p_stack stack)
{
	return get_stck_size(stack)==STACK_EMPTY;
}

_Bool stack_is_valid(p_stack stack)
{
	return stack!=NULL;
}

_Bool push(p_stack stack, stck_elem_t elem)
{
	if(stack_is_invalid(stack))
		return FALSE;

	p_snode snode=init_snode(elem);
	if(!snode)
		return FALSE;

	snode->next=stack->top;
	stack->top=snode;
	stack->size++;
	return TRUE;
}

_Bool pop(p_stack stack, stck_elem_t* elem)
{
	if(stack_is_invalid(stack))
		return FALSE;

	if(stack_is_empty(stack))
		return FALSE;

	p_snode pop_node=stack->top;
	stack->top=pop_node->next;
	*elem=pop_node->elem;

	free(pop_node);
	stack->size--;
	return TRUE;
}

p_stack del_stack(p_stack stack)
{
	if(stack_is_valid(stack))
	{
		stck_elem_t tmp;
		while(pop(stack,&tmp));
		free(stack);
		stack=NULL;
	}
	return stack;
}

