class T2 < Formula
    desc "A tui app made in go to track the time you spend on a projects"
    homepage "https://github.com/franpfeiffer/t2"
    url "https://github.com/franpfeiffer/t2/releases/tag/v0.1.0/app-linux-amd64" 

    version "0.1.0"
    sha256 "f3819c7c585084163ec665a49a0f2122b80f0b69b53a3c0521132453d09e3051"

    def install
      bin.install "app-linux-amd64" => "t2"
    end

    test do
      system "#{bin}/t2", "--version"
    end
end
