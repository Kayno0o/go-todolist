import { PriorityObjectType, PriorityObjects } from '../types/PriorityType';
import { JSONTask, Task } from '../types/TaskType';

export function JsonToTask(json: JSONTask): Task {
  const task = {
    name: json.name,
    task_group: json.task_group,
    priority: Object.values(PriorityObjects).find((priority) => priority.id === json.priority) as PriorityObjectType,
  };

  if (json.type === 'quantity') {
    return {
      ...task,
      type: json.type,
      amount: json.amount,
      completed: Number(json.completed),
    };
  } else {
    return {
      ...task,
      type: json.type,
      completed: json.completed === 1 ? true : false,
    };
  }
}
