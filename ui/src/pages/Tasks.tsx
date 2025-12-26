import { Page, PageTitle, Section, SectionTitle } from "@/components/atoms/Page"
import { Select } from "@/components/molecules/Select"
import { TaskHistory } from "@/components/task/TaskHistory"
import { TaskTable } from "@/components/task/TaskTable"
import { useTaskGetAll } from "@/lib/api/task"
import { TaskHistoryFilter, TaskResult } from "@/lib/types/task"
import { Chip, Group } from "@mantine/core"
import { useState } from "react"

export const Tasks = () => {
  const [filter, setFilter] = useState<TaskHistoryFilter>({});

  const { data: tasks, isLoading: isLoadingTasks } = useTaskGetAll()

  const handleTaskChange = (value: string | null) => {
    setFilter({ ...filter, uid: value ? value : undefined })
  }

  const handleResultChange = (value: string | string[]) => {
    let newResult: TaskResult | undefined = undefined

    if (value !== "all") newResult = value as TaskResult

    setFilter({ ...filter, result: newResult })
  }

  return (
    <Page>
      <PageTitle
        title="Background tasks"
        description="An overview of all background tasks."
      />

      <Section className="flex-none">
        <SectionTitle
          title="Active tasks"
          description="A refresh is recommened after a manually started task has finished."
        />
        <TaskTable
          tasks={tasks ?? []}
          isLoading={isLoadingTasks}
        />
      </Section>

      <Section className="min-h-full">
        <SectionTitle
          title="Run history"
          description="Recent runs across all background tasks."
        />
        <Group gap="xs">
          <Select
            data={tasks?.map(t => ({ value: t.uid, label: t.name }))}
            value={filter.uid}
            onChange={handleTaskChange}
            placeholder="Filter by task name..."
            disabled={isLoadingTasks}
          />
          <Chip.Group value={filter.result ? filter.result : "all"} onChange={handleResultChange}>
            <Chip value="all" color="secondary.2">All</Chip>
            <Chip value="success" color="secondary.2">Success</Chip>
            <Chip value="failed" color="secondary.2">Error</Chip>
          </Chip.Group>
        </Group>
        <TaskHistory filter={filter} />
      </Section>
    </Page>
  )
}
