class T2 < Formula
    desc "A tui app made in go to track the time you spend on a projects"
    homepage "https://github.com/franpfeiffer/t2"
    url "https://github.com/franpfeiffer/t2/releases/tag/0.1.0/t2-linux-amd64" 

    version "0.1.0"
    sha256 "36bcc1eb46fe07caa8621e4e338808b16f8a04edbfd72438bd351edace2faa24"

    def install
      bin.install "t2-linux-amd64" => "t2"
    end

    test do
      system "#{bin}/t2", "--version"
    end
end
