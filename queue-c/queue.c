#include<stdlib.h>
#include"queue.h"

_Bool queue_is_invalid(p_queue queue)
{
	return !queue_is_valid(queue);
};

p_qnode init_qnode(que_elem_t elem)
{
	p_qnode qnode=(p_qnode)malloc(sizeof(struct qnode));
	if(qnode)
	{
		qnode->elem=elem;
		qnode->next=NULL;
	}
	return qnode;
}

p_queue init_queue()
{
	p_queue queue=(p_queue)malloc(sizeof(struct queue));
	if(queue)
	{
		queue->head=NULL;
		queue->tail=NULL;
		queue->size=QUEUE_EMPTY;
	}
	return queue;
}

que_size_t get_que_size(p_queue queue)
{
	return queue->size;
}

_Bool queue_is_valid(p_queue queue)
{
	return queue!=NULL;
}

_Bool queue_is_empty(p_queue queue)
{
	return get_que_size(queue)==QUEUE_EMPTY;
}

_Bool en_queue(p_queue queue, que_elem_t elem)
{
	if(queue_is_invalid(queue))
		return FALSE;

	p_qnode qnode=init_qnode(elem);
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

_Bool de_queue(p_queue queue,que_elem_t* elem)
{
	if(queue_is_invalid(queue))
		return FALSE;

	if(queue_is_empty(queue))
		return FALSE;

	p_qnode de_node=queue->head;
	if(queue->head==queue->tail)
		queue->head=queue->tail=NULL;
	else
		queue->head=queue->head->next;

	*elem=de_node->elem;
	free(de_node);
	queue->size--;
	return TRUE;
}

p_queue del_queue(p_queue queue)
{
	if(queue_is_valid(queue))
	{
		que_elem_t tmp;
		while(de_queue(queue,&tmp));
		free(queue);
		queue=NULL;
	}
	return queue;
}

