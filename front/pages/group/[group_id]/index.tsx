import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useRouter } from 'next/router';
import React, { useEffect } from 'react';
import AddButton from '../../../app/components/AddButton';
import BackButton from '../../../app/components/BackButton';
import Layout from '../../../app/components/layout/Layout';
import PageTitle from '../../../app/components/PageTitle';
import TaskList from '../../../app/components/TaskList';
import { JSONGroup, TaskGroup } from '../../../app/types/GroupList';
import { TextColors } from '../../../app/utils/Colors';
import { JsonToGroup } from '../../../app/utils/Group';

export default function SingleGroup() {
  const router = useRouter();
  const groupId = router.query.group_id;
  const [group, setGroup] = React.useState<TaskGroup>();

  useEffect(() => {
    fetch(`/api/task_group/one/id/${groupId}`)
      .then((res) => res.json())
      .then(async (data: JSONGroup) => {
        const g = await JsonToGroup(data);
        setGroup(g);
      })
      .catch((err) => {
        console.log(err);
      });
  }, [groupId]);

  return (
    <Layout>
      {group && (
        <>
          <PageTitle className={TextColors[group.color]}>
            <BackButton route={`/`} />
            <FontAwesomeIcon icon={group.icon} className="mr-3" />
            {group.name}
          </PageTitle>

          <TaskList tasks={group.tasks} />

          <AddButton
            onClick={() => {
              router.push(`/group/${group.id}/task/new`);
            }}
          />
        </>
      )}
    </Layout>
  );
}
