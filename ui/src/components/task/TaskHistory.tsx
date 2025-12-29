import { useTaskGetHistory } from "@/lib/api/task";
import { TaskHistoryFilter } from "@/lib/types/task";
import { formatDate } from "@/lib/utils";
import { ActionIcon } from "@mantine/core";
import { useState } from "react";
import { LuTriangle } from "react-icons/lu";
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

  const [expandedIds, setExpandedIds] = useState<number[]>([])

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
          title: "",
          width: 40,
          textAlign: 'right',
          render: ({ id, error }) => error && (
            <ActionIcon variant="transparent" color="black" size="xs">
              <LuTriangle className={`fill-black w-2 ${expandedIds[0] === id && "rotate-180"}`} />
            </ActionIcon>
          )
        },
      ]}
      rowExpansion={{
        content: ({ record: { error } }) => (
          <div className="m-4">
            <p className="font-bold text-red-500">Error</p>
            <p className="text-red-500">{error}</p>
          </div>
        ),
        expandable: ({ record: { error } }) => Boolean(error),
        expanded: {
          recordIds: expandedIds,
          onRecordIdsChange: setExpandedIds,
        },
      }}
      records={history}
      noRecordsText="No tasks run yet"
      fetching={isLoading}
      onScrollToBottom={handleBottom}
    />
  )
}

