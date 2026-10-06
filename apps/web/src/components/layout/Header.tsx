import { useState } from "react";
import {
  AppBar,
  Box,
  Button,
  IconButton,
  Menu,
  MenuItem,
  Toolbar,
  Typography,
} from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "../../app/providers/useAuth";

export function Header() {
  const { isAuthenticated, user, logout } = useAuth();
  const navigate = useNavigate();
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);

  const closeMenu = () => setAnchorEl(null);

  const handleLogout = () => {
    closeMenu();
    logout();
    navigate("/login");
  };

  const links = [
    { label: "みんなの旅行記", to: "/public/travels" },
    ...(isAuthenticated ? [{ label: "自分の旅行記", to: "/travels" }] : []),
  ];

  return (
    <AppBar className="Header" position="sticky" elevation={0}>
      <Toolbar>
        <Typography variant="h6" component="div" sx={{ mr: 2 }}>
          Backpacker
        </Typography>

        {/* sm 以上：ボタン表示 */}
        <Box sx={{ display: { xs: "none", md: "flex" }, flexGrow: 1 }}>
          {links.map((link) => (
            <Button key={link.to} color="inherit" component={Link} to={link.to}>
              {link.label}
            </Button>
          ))}
        </Box>

        <Box sx={{ display: { xs: "none", md: "flex" }, alignItems: "center" }}>
          {isAuthenticated ? (
            <>
              <Typography sx={{ mr: 2 }}>Hello, {user?.nickname}</Typography>
              <Button color="inherit" onClick={handleLogout}>
                Logout
              </Button>
            </>
          ) : (
            <Button color="inherit" component={Link} to="/login">
              ログイン
            </Button>
          )}
        </Box>

        {/* xs：ハンバーガー */}
        <Box sx={{ display: { xs: "flex", md: "none" }, ml: "auto" }}>
          <IconButton
            color="inherit"
            aria-label="メニューを開く"
            onClick={(e) => setAnchorEl(e.currentTarget)}
          >
            <MenuIcon />
          </IconButton>
          <Menu anchorEl={anchorEl} open={Boolean(anchorEl)} onClose={closeMenu}>
            {links.map((link) => (
              <MenuItem
                key={link.to}
                component={Link}
                to={link.to}
                sx={{ py: 2 }}
                onClick={closeMenu}
              >
                {link.label}
              </MenuItem>
            ))}
            {isAuthenticated ? (
              <MenuItem sx={{ py: 2 }} onClick={handleLogout}>ログアウト</MenuItem>
            ) : (
              <MenuItem component={Link} to="/login" sx={{ py: 2 }} onClick={closeMenu}>
                ログイン
              </MenuItem>
            )}
          </Menu>
        </Box>
      </Toolbar>
    </AppBar>
  );
}
