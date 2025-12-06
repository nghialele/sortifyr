import { User } from "@/lib/types/user";
import { Avatar as MAvatar } from "@mantine/core";

type Props = {
  user: User | null;
}

export const Avatar = ({ user }: Props) => {
  if (user?.hasProfile) return <MAvatar src={`/api/user/image/${user.id}`} />

  return <MAvatar name={user?.name} color="black" />
}
