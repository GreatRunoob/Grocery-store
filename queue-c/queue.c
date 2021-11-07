#include <stdlib.h>
#include "queue.h"

static bool queue_is_invalid(queue_ptr queue)
{
	return !queue_is_valid(queue);
};

static bool queue_is_empty(queue_ptr queue)
{
	return get_queue_size(queue) == QUEUE_EMPTY;
}

static qnode_ptr init_qnode(queue_elem_t elem)
{
	qnode_ptr qnode = (qnode_ptr)malloc(sizeof(struct qnode));
	if (qnode)
	{
		qnode->elem = elem;
		qnode->next = NULL;
	}
	return qnode;
}

queue_ptr init_queue()
{
	queue_ptr queue = (queue_ptr)malloc(sizeof(struct queue));
	if (queue)
	{
		queue->head = NULL;
		queue->tail = NULL;
		queue->size = QUEUE_EMPTY;
	}
	return queue;
}

bool en_queue(queue_ptr queue, queue_elem_t elem)
{
	if (queue_is_invalid(queue))
		return false;

	qnode_ptr qnode = init_qnode(elem);
	if (!qnode)
		return false;

	if (queue_is_empty(queue))
	{
		queue->head = qnode;
		queue->tail = qnode;
	}
	else
	{
		queue->tail->next = qnode;
		queue->tail = qnode;
	}
	queue->size++;
	return true;
}

bool de_queue(queue_ptr queue, queue_elem_t *elem)
{
	if (queue_is_invalid(queue))
		return false;

	if (queue_is_empty(queue))
		return false;

	qnode_ptr de_node = queue->head;
	if (queue->head == queue->tail)
		queue->head = queue->tail = NULL;
	else
		queue->head = queue->head->next;

	*elem = de_node->elem;
	free(de_node);
	queue->size--;
	return true;
}

bool empty_queue(queue_ptr queue)
{
	if (queue_is_invalid(queue))
		return false;

	queue_elem_t tmp;
	while (de_queue(queue,&tmp));
	return true;
}

queue_ptr del_queue(queue_ptr queue)
{
	if (empty_queue(queue))
	{
		free(queue);
		queue = NULL;
	}
	return queue;
}

