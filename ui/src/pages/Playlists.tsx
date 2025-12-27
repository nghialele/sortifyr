import { Page, PageTitle, Section } from "@/components/atoms/Page"
import { Segment } from "@/components/molecules/Segment"
import { PlaylistDuplicates } from "@/components/playlist/PlaylistDuplicates"
import { PlaylistOverview } from "@/components/playlist/PlaylistOverview"
import { PlaylistUnplayables } from "@/components/playlist/PlaylistUnplayable"
import { Center, Group } from "@mantine/core"
import { ReactNode, useMemo, useState } from "react"
import { LuListX, LuSquareStack, LuTextSearch } from "react-icons/lu"

type ViewOption = "overview" | "duplicates" | "unplayables"
type View = { value: ViewOption, label: string, icon: ReactNode, }

const views: View[] = [
  {
    value: "overview",
    label: "Overview",
    icon: <LuTextSearch />,
  },
  {
    value: "duplicates",
    label: "Duplicates",
    icon: <LuSquareStack />,
  },
  {
    value: "unplayables",
    label: "Unplayable",
    icon: <LuListX />,
  },
]

const renderView = (view: ViewOption) => {
  switch (view) {
    case "overview": return <PlaylistOverview />
    case "duplicates": return <PlaylistDuplicates />
    case "unplayables": return <PlaylistUnplayables />
  }
}

const storageKey = "sortifyr-playlist-view"

export const Playlists = () => {
  const [view, setView] = useState<ViewOption>(localStorage.getItem(storageKey) as ViewOption ?? "overview")

  const handleSegment = (view: ViewOption) => {
    localStorage.setItem(storageKey, view)
    setView(view)
  }

  const renderedView = useMemo(() => renderView(view), [view])

  return (
    <Page>
      <Group justify="space-between">
        <PageTitle
          title="Playlists"
          description="An overview of your Spotify playlists."
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
        />
      </Group>

      <Section>
        {renderedView}
      </Section>

    </Page>
  )
}
