import { TaskResult as TaskResultType } from "@/lib/types/task"
import { capitalize } from "@/lib/utils"
import { Group } from "@mantine/core"

type Props = {
  result: TaskResultType
}

export const TaskResult = ({ result }: Props) => {
  return (
    <Group gap="xs">
      <div className={`w-2 h-2 rounded-full ${result === TaskResultType.Success ? "bg-green-500" : "bg-red-500"}`} />
      <p>{capitalize(result)}</p>
    </Group>
  )
}
