#ifndef _CQUEUE_H_
#define _CQUEUE_H_

typedef unsigned int que_elem_t;
typedef unsigned int que_size_t;

typedef struct qnode
{
	que_elem_t elem;
	struct qnode* next;
} *qnode_ptr;

typedef struct queue
{
	qnode_ptr head;
	qnode_ptr tail;
	que_size_t size;
} *queue_ptr;

qnode_ptr init_qnode(que_elem_t elem);

queue_ptr init_queue();

que_size_t get_que_size(queue_ptr queue);

_Bool queue_is_empty(queue_ptr queue);

_Bool queue_is_valid(queue_ptr queue);

_Bool en_queue(queue_ptr queue, que_elem_t elem);

_Bool de_queue(queue_ptr queue, que_elem_t* elem);

_Bool empty_queue(queue_ptr queue);

queue_ptr del_queue(queue_ptr queue);

#define QUEUE_EMPTY (que_size_t) 0
#define FALSE (_Bool) 0
#define TRUE (_Bool) 1

#endif
