import { useRouter } from 'next/router';
import React from 'react';
import BackButton from '../../../../app/components/BackButton';
import Input from '../../../../app/components/form/Input';
import Select from '../../../../app/components/form/Select';
import Layout from '../../../../app/components/layout/Layout';
import PageTitle from '../../../../app/components/PageTitle';
import TaskCard from '../../../../app/components/TaskCard';
import { PriorityObjects, PriorityType } from '../../../../app/types/PriorityType';
import { TaskType } from '../../../../app/types/TaskType';

export default function NewGroup() {
  const [name, setName] = React.useState('');
  const [type, setType] = React.useState<TaskType>('tick');
  const [amount, setAmount] = React.useState<number>(0);
  const [priority, setPriority] = React.useState<PriorityType>('low');
  const [error, setError] = React.useState<string>('');

  const router = useRouter();
  const groupId = Number(router.query.group_id);

  const groupCard = React.useMemo(() => {
    if (type === 'tick') {
      return (
        <TaskCard
          className="mt-4"
          task={{ name, type, priority: PriorityObjects[priority], task_group: groupId, completed: false }}
        />
      );
    } else {
      return (
        <TaskCard
          className="mt-4"
          task={{ name, type, amount, priority: PriorityObjects[priority], task_group: groupId, completed: 0 }}
        />
      );
    }
  }, [name, type, amount, priority, groupId]);

  return (
    <Layout>
      <PageTitle>
        <BackButton route={`/group/${groupId}`} />
        New Task
      </PageTitle>

      <form
        method="POST"
        className="flex flex-col gap-2"
        onSubmit={(e) => {
          e.preventDefault();

          if (name.length < 1) return setError('Name is required');
          if (type === 'quantity' && amount < 1) return setError('Amount is required');

          fetch('/api/task/new', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              amount: type === 'tick' ? null : amount,
              completed: type === 'tick' ? false : 0,
              priority: PriorityObjects[priority].id,
              task_group: groupId,
              name,
              type,
            }),
          })
            .then((res) => res.json())
            .then((data) => {
              if (data.status === '500') {
                return setError(data.message);
              }

              router.push(`/group/${groupId}`);
            });
        }}
      >
        <Input name="name" label="Name" value={name} onChange={(e) => setName(e.target.value)} />

        <Select
          name="type"
          label="Type"
          value={type}
          onChange={(e) => setType(e.target.value as TaskType)}
          options={[
            { value: 'tick', label: 'Tick' },
            { value: 'quantity', label: 'Quantity' },
          ]}
        />

        {type === 'quantity' && (
          <Input
            name="amount"
            label="Amount to reach"
            type="number"
            value={String(amount)}
            onChange={(e) => setAmount(Number(e.target.value))}
          />
        )}

        <Select
          name="priority"
          label="Priority"
          value={priority}
          onChange={(e) => setPriority(e.target.value as PriorityType)}
          options={[
            { value: 'low', label: 'Low' },
            { value: 'medium', label: 'Medium' },
            { value: 'high', label: 'High' },
          ]}
        />

        {error && <div className="text-red-500">{error}</div>}

        <div className="flex justify-end mt-2">
          <button type="submit" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg">
            Create
          </button>
        </div>
      </form>

      <div className="mt-4">
        <h2 className="text-xl">Preview:</h2>
        {groupCard}
      </div>
    </Layout>
  );
}
