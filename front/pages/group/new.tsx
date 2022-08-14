import { useRouter } from 'next/router';
import React from 'react';
import BackButton from '../../app/components/BackButton';
import Input from '../../app/components/form/Input';
import Select from '../../app/components/form/Select';
import GroupCard from '../../app/components/GroupCard';
import Layout from '../../app/components/layout/Layout';
import PageTitle from '../../app/components/PageTitle';
import { BrightBackgroundColors, ColorType } from '../../app/utils/Colors';
import { IconType, TaskIcons } from '../../app/utils/Icons';
import { FirstUpper } from '../../app/utils/Utils';

export default function NewGroup() {
  const [name, setName] = React.useState('');
  const [description, setDescription] = React.useState<string>('');
  const [color, setColor] = React.useState<ColorType>('purple');
  const [icon, setIcon] = React.useState<IconType>('Headset');
  const [error, setError] = React.useState<string>('');

  const groupCard = React.useMemo(() => {
    return <GroupCard className="mt-4" group={{ name, description, color, icon: TaskIcons[icon], tasks: [] }} />;
  }, [name, description, color, icon]);

  const router = useRouter();

  return (
    <Layout>
      <PageTitle>
        <BackButton route={`/`} />
        New Group
      </PageTitle>
      <form
        method="POST"
        className="flex flex-col gap-2"
        onSubmit={(e) => {
          e.preventDefault();

          if (name.length < 1) return setError('Name is required');

          fetch('/api/task_group/new', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name, description, color, icon }),
          })
            .then((res) => res.json())
            .then((data) => {
              router.push(`/group/${data.id}`);
            });
        }}
      >
        <Input name="name" label="Name" value={name} onChange={(e) => setName(e.target.value)} />

        <Input
          name="description"
          label="Description"
          type="textarea"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />

        <Select
          label="Color"
          name="color"
          value={color}
          onChange={(e) => setColor(e.target.value as ColorType)}
          options={Object.keys(BrightBackgroundColors).map((key) => {
            return {
              value: key,
              label: FirstUpper(key),
            };
          })}
        />

        <Select
          label="Icon"
          name="icon"
          value={icon}
          onChange={(e) => setIcon(e.target.value as IconType)}
          options={Object.keys(TaskIcons).map((key) => {
            return { value: key, label: key };
          })}
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
