require 'yaml'

def rubdom
  members = YAML.load_file("./members.yml")['rubdom']['members']
  members[rand(members.size)]
end

def progress_bar(i, max = 100)
  i = max if i > max
  rest_size = 1 + 5 + 1
  bar_width = 79 - rest_size
  percent = i * 100.0 / max
  bar_length = i * bar_width.to_f / max
  bar_str = ('#' * bar_length).ljust(bar_width)
  progress_num = '%3.1f' % percent
  print "\r#{bar_str} #{'%5s' % progress_num}%"
end

puts "選ばれし者を探しています..."

(0..1000).each do |j|
  sleep 0.001
  progress_bar(j, 1000)
end

puts "\n選ばれし者は#{rubdom}でした"
