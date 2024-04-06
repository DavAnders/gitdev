
export default function GithubStats({ user }) {
    return (
      <div className="flex space-x-4">
        <div className="flex flex-col items-center">
          <div className="text-2xl font-bold">{user.Repos.length}</div>
          <div className="text-sm">Repos</div>
        </div>
        <div className="flex flex-col items-center">
          <div className="text-2xl font-bold">{user.Profile.Followers}</div>
          <div className="text-sm">Followers</div>
        </div>
        <div className="flex flex-col items-center">
          <div className="text-2xl font-bold">{user.Profile.Following}</div>
          <div className="text-sm">Following</div>
        </div>
      </div>
    );
  }