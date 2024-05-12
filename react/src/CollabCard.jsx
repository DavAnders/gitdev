import Avatar from "./Avatar";

export default function CollabCard({collab}) {
    return(
        <div className="p-3 border-b border-gray-600">
            <div className="flex items-center space-x-1">
                <Avatar src={collab.AvatarUrl} size="20px" />
                <div>
                    <h1>{collab.Username}</h1>
                    <h1>{collab.Title}</h1>
                </div>
                
            </div>
            <h2>{collab.Email}</h2>
            <h2>{collab.Location}</h2>
        </div>
    );
};
