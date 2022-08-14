import React, { useEffect, useState } from 'react';
import type { NextPage } from 'next';
import Layout from '../app/components/layout/Layout';
import GroupCard from '../app/components/GroupCard';
import { JSONGroup, TaskGroup } from '../app/types/GroupList';
import { JsonToGroup } from '../app/utils/Group';
import { useRouter } from 'next/router';
import PageTitle from '../app/components/PageTitle';
import AddButton from '../app/components/AddButton';

const Home: NextPage = () => {
  const [groups, setGroups] = useState<TaskGroup[]>([]);

  useEffect(() => {
    fetch(`/api/task_group/all`)
      .then((res) => res.json())
      .then(async (data: JSONGroup[]) => {
        const groups: TaskGroup[] = [];

        for (const group of data) {
          groups.push(await JsonToGroup(group));
        }

        setGroups(groups);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  const router = useRouter();

  return (
    <Layout>
      <PageTitle>Lists</PageTitle>

      <div className="grid grid-cols-[repeat(auto-fit,_minmax(150px,_1fr))] gap-4">
        {groups.map((group, index) => (
          <GroupCard
            key={index}
            group={group}
            onClick={() => {
              router.push(`/group/${group.id}`);
            }}
          />
        ))}
      </div>

      <AddButton
        onClick={() => {
          router.push('/group/new');
        }}
      />
    </Layout>
  );
};

export default Home;
