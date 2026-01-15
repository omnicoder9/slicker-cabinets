# slicker-cabinets

ubuntu:
sudo apt install tmux
echo -e "set-option -g history-limit 100000\nset -g mouse on" > ~/.tmux.conf
OR if tmux already running:
echo -e "set-option -g history-limit 100000\nset -g mouse on" > ~/.tmux.conf && tmux source-file ~/.tmux.conf

npm install -g @google/gemini-cli
gemini --version

gemini --model gemini-3-pro-preview

enter gemini api key

go run main.go

localhost:8081