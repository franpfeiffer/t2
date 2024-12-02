class T2 < Formula
    desc "A tui app made in go to track the time you spend on a projects"
    homepage "https://github.com/franpfeiffer/t2"
    url "https://github.com/franpfeiffer/t2/releases/tag/0.1.0/t2-linux-amd64" 

    version "0.1.0"
    sha256 "2321c291093612f525576bf9002038377f46e690058b3f39923a64bc27972aef"

    def install
      bin.install "t2-linux-amd64" => "t2"
    end

    test do
      system "#{bin}/t2", "--version"
    end
end
