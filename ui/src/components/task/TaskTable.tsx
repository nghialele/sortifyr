import { useTaskStart } from "@/lib/api/task";
import { Task, TaskStatus } from "@/lib/types/task";
import { ActionIcon } from "@mantine/core";
import { notifications } from "@mantine/notifications";
import { formatDistanceToNow, formatDuration, intervalToDuration } from "date-fns";
import { LuCirclePlay } from "react-icons/lu";
import { Table } from "../molecules/Table";
import { TaskResult } from "./TaskResult";

type Props = {
  tasks: Task[];
  isLoading: boolean;
}

const formatInterval = (ms: number) => {
  const duration = intervalToDuration({ start: 0, end: ms / 1000000 });
  return formatDuration(duration);
}

export const TaskTable = ({ tasks, isLoading }: Props) => {
  const taskRun = useTaskStart()

  const handleClick = (task: Task) => {
    taskRun.mutate(task, {
      onSuccess: () => notifications.show({ variant: "success", title: task.name, message: "Started" })
    })
  }

  return (
    <Table
      idAccessor="uid"
      columns={[
        {
          accessor: "uid",
          title: "",
          width: 38,
          render: task => (
            <ActionIcon onClick={() => handleClick(task)} variant="transparent" color="secondary.3" loading={task.status === TaskStatus.Running}>
              <LuCirclePlay />
            </ActionIcon>
          )
        },
        { accessor: "name", title: "Task" },
        { accessor: "lastRun", render: ({ lastRun }) => <p className="text-muted">{lastRun ? formatDistanceToNow(lastRun, { addSuffix: true }) : ""}</p> },
        { accessor: "nextRun", render: ({ nextRun }) => <p className="text-muted">{formatDistanceToNow(nextRun, { addSuffix: true })}</p> },
        { accessor: "interval", render: ({ interval }) => <p className="text-muted">{interval ? formatInterval(interval) : ""}</p> },
        {
          accessor: "lastStatus",
          title: "Last result",
          render: ({ lastStatus }) => lastStatus && <TaskResult result={lastStatus} />
        },
      ]}
      records={tasks}
      fetching={isLoading}
    />
  )
}
