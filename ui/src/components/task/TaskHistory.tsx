import { useTaskGetHistory } from "@/lib/api/task";
import { TaskHistoryFilter } from "@/lib/types/task";
import { formatDate } from "@/lib/utils";
import { Table } from "../molecules/Table";
import { TaskResult } from "./TaskResult";

type Props = {
  filter?: TaskHistoryFilter;
}

const formatDuration = (nanos: number) => {
  const ms = Math.floor(nanos / 1_000_000) % 1000
  const s = Math.floor(nanos / 1_000_000_000)

  const msString = `${ms.toString().padStart(3, '0')}ms`
  const sString = `${s.toString().padStart(2, '0')}s `

  return <span>{sString}<span className="text-muted-foreground">{msString}</span></span>
}

export const TaskHistory = ({ filter }: Props) => {
  const { history, isLoading, isFetchingNextPage, hasNextPage, fetchNextPage } = useTaskGetHistory(filter)

  const handleBottom = () => {
    if (!hasNextPage) return
    if (isFetchingNextPage) return

    fetchNextPage()
  }

  return (
    <Table
      columns={[
        { accessor: "name", title: "Task" },
        { accessor: "runAt", render: ({ runAt }) => <p className="text-muted">{formatDate(runAt)}</p> },
        { accessor: "duration", render: ({ duration }) => <p className="text-muted">{formatDuration(duration)}</p> },
        {
          accessor: "result",
          title: "Result",
          render: ({ result }) => <TaskResult result={result} />
        },
        {
          accessor: "error",
          title: "Message",
          render: task => <p className="text-muted">{task.error ? task.error : task.message}</p>
        },
      ]}
      records={history}
      noRecordsText="No tasks run yet"
      fetching={isLoading}
      onScrollToBottom={handleBottom}
    />
  )
}

