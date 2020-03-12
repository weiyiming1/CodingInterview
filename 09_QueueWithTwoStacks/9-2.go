package main

type TwoQueueStack struct {
	queue1 []interface{}
	queue2 []interface{}
}

func NewTwoQueueStack() TwoQueueStack {
	return TwoQueueStack{}
}

func (t *TwoQueueStack) Push(val interface{}) {
	if len(t.queue1) == 0 && len(t.queue2) == 0 {
		t.queue1 = append(t.queue1, val) // 两个都为空，默认选择queue1
	} else if len(t.queue1) != 0 && len(t.queue2) == 0 {
		t.queue1 = append(t.queue1, val)
	} else if len(t.queue1) == 0 && len(t.queue2) != 0 {
		t.queue2 = append(t.queue2, val)
	}
}

func (t *TwoQueueStack) Pop() interface{} {
	if t.queue1 == nil && t.queue2 == nil {
		return nil
	}
	var res interface{}
	if len(t.queue1) == 0 && len(t.queue2) != 0 {
		for len(t.queue2) > 0 {
			res = t.queue2[0]
			t.queue2 = append(t.queue2[1:])
			if len(t.queue2) != 0 {
				t.queue1 = append(t.queue1, res)
			}
		}
	} else if len(t.queue1) != 0 && len(t.queue2) == 0 {
		for len(t.queue1) > 0 {
			res = t.queue1[0]
			t.queue1 = append(t.queue1[1:])
			if len(t.queue1) != 0 {
				t.queue2 = append(t.queue2, res)
			}
		}
	}
	return res
}
