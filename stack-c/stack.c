#include<stdlib.h>
#include"stack.h"

static bool stack_is_invalid(stack_ptr stack)
{
	return !stack_is_valid(stack);
};

static bool stack_is_empty(stack_ptr stack)
{
	return get_stack_size(stack) == STACK_EMPTY;
}

static snode_ptr init_snode(stack_elem_t elem)
{
	snode_ptr snode = (snode_ptr)malloc(sizeof(struct snode));
	if (snode)
	{
		snode->elem = elem;
		snode->next = NULL;
	}
	return snode;
}

stack_ptr init_stack()
{
	stack_ptr stack = (stack_ptr)malloc(sizeof(struct stack));
	if (stack)
	{
		stack->top = NULL;
		stack->size = STACK_EMPTY;
	}
	return stack;
}

bool push(stack_ptr stack, stack_elem_t elem)
{
	if (stack_is_invalid(stack))
		return false;

	snode_ptr snode = init_snode(elem);
	if (!snode)
		return false;

	snode->next = stack->top;
	stack->top = snode;
	stack->size++;
	return true;
}

bool pop(stack_ptr stack, stack_elem_t *elem)
{
	if (stack_is_invalid(stack))
		return false;

	if (stack_is_empty(stack))
		return false;

	snode_ptr pop_node = stack->top;
	stack->top = pop_node->next;
	*elem = pop_node->elem;

	free(pop_node);
	stack->size--;
	return true;
}

bool empty_stack(stack_ptr stack)
{
	if (stack_is_invalid(stack))
		return false;

	stack_elem_t tmp;
	while (pop(stack,&tmp));
	return true;
}

stack_ptr del_stack(stack_ptr stack)
{
	if (empty_stack(stack))
	{
		free(stack);
		stack = NULL;
	}
	return stack;
}

