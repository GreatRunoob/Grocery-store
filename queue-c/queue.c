#include<stdlib.h>
#include"queue.h"

_Bool queue_is_invalid(queue_ptr queue)
{
	return !queue_is_valid(queue);
};

qnode_ptr init_qnode(que_elem_t elem)
{
	qnode_ptr qnode=(qnode_ptr)malloc(sizeof(struct qnode));
	if(qnode)
	{
		qnode->elem=elem;
		qnode->next=NULL;
	}
	return qnode;
}

queue_ptr init_queue()
{
	queue_ptr queue=(queue_ptr)malloc(sizeof(struct queue));
	if(queue)
	{
		queue->head=NULL;
		queue->tail=NULL;
		queue->size=QUEUE_EMPTY;
	}
	return queue;
}

que_size_t get_que_size(queue_ptr queue)
{
	return queue->size;
}

_Bool queue_is_valid(queue_ptr queue)
{
	return queue!=NULL;
}

_Bool queue_is_empty(queue_ptr queue)
{
	return get_que_size(queue)==QUEUE_EMPTY;
}

_Bool en_queue(queue_ptr queue, que_elem_t elem)
{
	if(queue_is_invalid(queue))
		return FALSE;

	qnode_ptr qnode=init_qnode(elem);
	if(!qnode)
		return FALSE;

	if(queue_is_empty(queue))
	{
		queue->head=qnode;
		queue->tail=qnode;
	}
	else
	{
		queue->tail->next=qnode;
		queue->tail=qnode;
	}
	queue->size++;
	return TRUE;
}

_Bool de_queue(queue_ptr queue,que_elem_t* elem)
{
	if(queue_is_invalid(queue))
		return FALSE;

	if(queue_is_empty(queue))
		return FALSE;

	qnode_ptr de_node=queue->head;
	if(queue->head==queue->tail)
		queue->head=queue->tail=NULL;
	else
		queue->head=queue->head->next;

	*elem=de_node->elem;
	free(de_node);
	queue->size--;
	return TRUE;
}

_Bool empty_queue(queue_ptr queue)
{
	if(queue_is_invalid(queue))
		return FALSE;

	que_elem_t tmp;
	while(de_queue(queue,&tmp));
	return TRUE;
}

queue_ptr del_queue(queue_ptr queue)
{
	if(empty_queue(queue))
	{
		free(queue);
		queue=NULL;
	}
	return queue;
}

