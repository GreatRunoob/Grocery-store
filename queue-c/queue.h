#ifndef _CQUEUE_H_
#define _CQUEUE_H_

typedef unsigned int que_elem_t;
typedef unsigned int que_size_t;

typedef struct qnode
{
	que_elem_t elem;
	struct qnode* next;
} *p_qnode;

typedef struct queue
{
	p_qnode head;
	p_qnode tail;
	que_size_t size;
} *p_queue;

p_qnode init_qnode(que_elem_t elem);

p_queue init_queue();

que_size_t get_que_size(p_queue queue);

_Bool queue_is_empty(p_queue queue);

_Bool queue_is_valid(p_queue queue);

_Bool en_queue(p_queue queue, que_elem_t elem);

_Bool de_queue(p_queue queue, que_elem_t* elem);

_Bool empty_queue(p_queue queue);

p_queue del_queue(p_queue queue);

#define QUEUE_EMPTY (que_size_t) 0
#define FALSE (_Bool) 0
#define TRUE (_Bool) 1

#endif
