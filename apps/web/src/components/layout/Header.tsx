import { Button, AppBar, Toolbar, Typography } from "@mui/material";
import { useNavigate } from "react-router-dom";

import { useAuth } from "../../app/providers/AuthProvider";

export function Header() {
  const { isAuthenticated, user, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate("/login");
  }

  return (
    <AppBar position="static" elevation={0}>
      <Toolbar>
        <Typography variant="h6" component="div">
          Backpacker
        </Typography>

        {isAuthenticated && (
          <>
            <Typography sx={{ ml: "auto", mr: 2 }}>
              Hello, {user?.nickname}
            </Typography>

            <Button color="inherit" onClick={handleLogout}>
              Logout
            </Button>
          </>
        )}
      </Toolbar>
    </AppBar>
  );
}
