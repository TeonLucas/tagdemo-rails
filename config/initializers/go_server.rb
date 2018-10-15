require 'quartz'

class Goserver
  GO_EXECUTABLE = 'goserver'

  def self.initialize
    @quartz = Quartz::Client.new(bin_path: Rails.root.join('bin', GO_EXECUTABLE).to_s)
  end

  def self.client
    return @quartz
  end
end

Goserver.initialize
