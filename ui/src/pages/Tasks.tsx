import { Page, PageTitle, Section, SectionTitle } from "@/components/atoms/Page"
import { Segment } from "@/components/molecules/Segment"
import { Select } from "@/components/molecules/Select"
import { TaskHistory } from "@/components/task/TaskHistory"
import { TaskTable } from "@/components/task/TaskTable"
import { useTaskGetAll } from "@/lib/api/task"
import { TaskHistoryFilter, TaskResult } from "@/lib/types/task"
import { capitalize } from "@/lib/utils"
import { Group } from "@mantine/core"
import { useState } from "react"

export const Tasks = () => {
  const [filter, setFilter] = useState<TaskHistoryFilter>({});

  const { data: tasks, isLoading: isLoadingTasks } = useTaskGetAll()

  const handleTaskChange = (value: string | null) => {
    setFilter({ ...filter, uid: value ? value : undefined })
  }

  const handleResultChange = (value: string) => {
    let result: TaskResult | undefined = undefined

    if (value !== "all") result = value as TaskResult

    setFilter({ ...filter, result: result })
  }

  const handleRecurringChange = (value: string) => {
    let recurring: boolean | undefined = undefined

    if (value != "all") recurring = value === "true"

    setFilter({ ...filter, recurring })
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
          <Segment
            data={[
              { value: "all", label: "All" },
              ...Object.values(TaskResult).map(r => ({ value: r, label: capitalize(r) }))
            ]}
            value={filter.result ? filter.result : "all"}
            onChange={handleResultChange}
            secondary
          />
          <Segment
            data={[
              { value: "all", label: "All" },
              { value: "true", label: "Recurring" },
              { value: "false", label: "Non-recurring" },
            ]}
            value={filter.recurring !== undefined ? String(filter.recurring) : "all"}
            onChange={handleRecurringChange}
            secondary
          />
        </Group>
        <TaskHistory filter={filter} />
      </Section>
    </Page>
  )
}
