#ifndef _CQUEUE_H_
#define _CQUEUE_H_

#include <stdbool.h>
#include <stddef.h>

#define QUEUE_EMPTY (queue_size_t) 0

typedef unsigned int queue_elem_t;
typedef unsigned int queue_size_t;

typedef struct qnode
{
	queue_elem_t elem;
	struct qnode *next;
} *qnode_ptr;

typedef struct queue
{
	qnode_ptr head;
	qnode_ptr tail;
	queue_size_t size;
} *queue_ptr;

static inline queue_size_t get_queue_size(queue_ptr queue)
{
	return queue->size;
}

static inline bool queue_is_valid(queue_ptr queue)
{
	return queue != NULL;
}

static bool queue_is_empty(queue_ptr queue);

static qnode_ptr init_qnode(queue_elem_t elem);

queue_ptr init_queue();

bool en_queue(queue_ptr queue, queue_elem_t elem);

bool de_queue(queue_ptr queue, queue_elem_t *elem);

bool empty_queue(queue_ptr queue);

queue_ptr del_queue(queue_ptr queue);

#endif
