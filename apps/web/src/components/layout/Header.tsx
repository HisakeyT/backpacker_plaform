import { AppBar, Toolbar, Typography } from "@mui/material";

export function Header() {
  return (
    <AppBar position="static" elevation={0}>
      <Toolbar>
        <Typography variant="h6" component="div">
          Backpacker
        </Typography>
      </Toolbar>
    </AppBar>
  );
}
