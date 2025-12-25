import { Group, Center, Pill } from "@mantine/core"
import { formatDistanceToNow } from "date-fns"
import { SectionTitle } from "../atoms/Page"
import { DirectoryTree } from "../directory/DirectoryTree"
import { Search } from "../molecules/Search"
import { Segment } from "../molecules/Segment"
import { PlaylistTable } from "./PlaylistTable"
import { useDirectoryGetAll } from "@/lib/api/directory"
import { usePlaylistGetAll } from "@/lib/api/playlist"
import { useTaskGetAll } from "@/lib/api/task"
import { convertDirectorySchema, Directory } from "@/lib/types/directory"
import { useDebouncedValue } from "@mantine/hooks"
import { useState, useMemo, ReactNode } from "react"
import { LuTable, LuFolderTree } from "react-icons/lu"

type ViewOption = "table" | "tree"
type View = { value: ViewOption, label: string, icon: ReactNode }

const views: View[] = [
  {
    value: "table",
    label: "Table",
    icon: <LuTable />,
  },
  {
    value: "tree",
    label: "Tree",
    icon: <LuFolderTree />,
  },
]

const storageKey = "sortifyr-playlist-overview-view"

const directoryHasPlaylist = (directory: Directory, search: string): boolean => {
  if (directory.playlists.find(p => p.name.toLowerCase().includes(search.toLowerCase()))) return true

  if (!directory.children) return false

  return directory.children.some(c => directoryHasPlaylist(c, search))
}

export const PlaylistOverview = () => {
  const [filter, setFilter] = useState("")
  const [debounced] = useDebouncedValue(filter, 200);
  const [view, setView] = useState<ViewOption>(localStorage.getItem(storageKey) as ViewOption ?? "table")

  const handleSegment = (view: ViewOption) => {
    localStorage.setItem(storageKey, view)
    setView(view)
  }

  const { data: playlistsAll, isLoading: isLoadingPlaylists } = usePlaylistGetAll()
  const { data: directoriesAll, isLoading: isLoadingDirectories } = useDirectoryGetAll()
  const { data: tasks } = useTaskGetAll()

  const playlists = useMemo(() => {
    return playlistsAll?.filter(p =>
      p.name.toLowerCase().includes(debounced.toLowerCase())
    ) ?? []
  }, [playlistsAll, debounced])
  const trackAmount = playlists.map(p => p.trackAmount).reduce((acc, curr) => acc + curr, 0)

  const directories = useMemo(() => {
    return convertDirectorySchema(directoriesAll?.filter(d => directoryHasPlaylist(d, debounced)) ?? [])
  }, [directoriesAll, debounced])

  const task = tasks?.find(t => t.uid === "task-playlist")

  return (
    <>
      <SectionTitle
        title="All playlists"
        description="Switch between a simple table and a directory tree."
      />
      <Group gap="xs">
        <Search
          placeholder="Filter by playlist name..."
          value={filter}
          onChange={e => setFilter(e.target.value)}
          className="grow"
        />
        <Segment
          data={views.map(v => ({
            value: v.value,
            label: (
              <Center style={{ gap: 4 }}>
                {v.icon}
                <p>{v.label}</p>
              </Center>
            )
          }))}
          value={view}
          onChange={e => handleSegment(e as ViewOption)}
          secondary
        />
      </Group>
      <Group gap="xs">
        <p className="text-sm text-muted">{`${playlists.length} playlists`}</p>
        <Pill bg="secondary.1">{`${trackAmount} tracks`}</Pill>
        {task && <p className="ml-auto text-sm text-muted">{`Next sync in ${formatDistanceToNow(task.nextRun)}`}</p>}
      </Group>
      {view === "table" ? (
        <PlaylistTable
          playlists={playlists}
          isLoading={isLoadingPlaylists}
        />
      ) : (
        <DirectoryTree
          roots={directories}
          isLoading={isLoadingDirectories}
          title="Playlists not in a directory are not shown"
        />
      )}
    </>
  )
}
